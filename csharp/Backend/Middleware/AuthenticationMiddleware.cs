using System.IdentityModel.Tokens.Jwt;
using Database;
using Database.Application;
using Database.Application.Models;

namespace Backend.Middleware;

public class AuthenticationMiddleware
{
    private readonly RequestDelegate _next;
    private readonly ILogger<AuthenticationMiddleware> _logger;
    private readonly DatabaseFactory<ApplicationDatabase> _databaseFactory;
    public AuthenticationMiddleware(RequestDelegate next, ILogger<AuthenticationMiddleware> logger, DatabaseFactory<ApplicationDatabase> databaseFactory)
    {
        _next = next;
        _logger = logger;
        _databaseFactory = databaseFactory;
    }

    public async Task InvokeAsync(HttpContext context)
    {
        var authorizationHeader = context.Request.Headers.Authorization.FirstOrDefault(x=> x is not null && x.StartsWith("Bearer "));
        
        if (authorizationHeader is null) 
        {
            context.Items["Role"] = Role.Guest();
            await _next(context);
            return;
        } 

        var token = authorizationHeader["Bearer ".Length..].Trim();
        
        using (var database = await _databaseFactory.New())
        {
            var validationResponse = await database.ValidateJwtToken(token);
            if (validationResponse is null || validationResponse.Session is null || validationResponse.Session.IsExpired()) 
            {
                context.Items["Role"] = Role.Guest();
                await _next(context);
                return;
            }
            context.Items["Role"] = validationResponse.Account.Role;
            context.Items["Account"] = validationResponse.Account;
            context.Items["Session"] = validationResponse.Session!;
        }
        
        await _next(context);
    }
}
