using Microsoft.AspNetCore.Authentication;
using Microsoft.AspNetCore.Authentication.Cookies;
using Microsoft.AspNetCore.Mvc;
using MySql.Data.MySqlClient;
using System.Security.Claims;

namespace Backend.Controllers
{
    [Route("api/auth")]
    [ApiController]
    public class AuthController(IConfiguration configuration) : ControllerBase
    {
        [HttpGet("github")]
        public IActionResult GitHubUserInfo(HttpContext ctx)
        {
            return new OkObjectResult(
                ctx.User.Claims.Select(x => new { x.Type, x.Value }).ToList()
            );
        }

        [HttpGet("github/login")]
        public IActionResult GitHubLogin()
        {
            var redirectUrl = Url.Action(nameof(GitHubCallback), "Auth", null, Request.Scheme);
            return Challenge(new AuthenticationProperties { RedirectUri = redirectUrl }, "GitHub");
        }

        [HttpGet("github/callback")]
        public async Task<IActionResult> GitHubCallback()
        {
            var authenticateResult = await HttpContext.AuthenticateAsync(CookieAuthenticationDefaults.AuthenticationScheme);
            if (!authenticateResult.Succeeded) return Unauthorized();

            var githubId = authenticateResult.Principal?.FindFirst("GitHubId")?.Value;
            var username = authenticateResult.Principal?.FindFirst(ClaimTypes.Name)?.Value;

            string sessionToken = Guid.NewGuid().ToString();
            using var connection = new MySqlConnection(configuration.GetConnectionString("DefaultConnection"));
            await connection.OpenAsync();

            string insertSessionQuery = @"
                INSERT INTO sessions (user_id, session_token, expires_at) 
                VALUES ((SELECT id FROM users WHERE github_id = @github_id), @session_token, DATE_ADD(NOW(), INTERVAL 1 DAY))";
            using var command = new MySqlCommand(insertSessionQuery, connection);
            command.Parameters.AddWithValue("@github_id", githubId);
            command.Parameters.AddWithValue("@session_token", sessionToken);
            await command.ExecuteNonQueryAsync();

            return Ok(new { message = $"Welcome, {username}!", session_token = sessionToken });
        }

        [HttpGet("logout")]
        public async Task<IActionResult> Logout()
        {
            await HttpContext.SignOutAsync(CookieAuthenticationDefaults.AuthenticationScheme);
            return Ok(new { message = "Logged out successfully" });
        }
    }
}
