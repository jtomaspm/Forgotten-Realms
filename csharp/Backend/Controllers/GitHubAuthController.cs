using Microsoft.AspNetCore.Authentication;
using Microsoft.AspNetCore.Mvc;

namespace Backend.Controllers
{
    [Route("api/auth/github")]
    [ApiController]
    public class GitHubAuthController(IConfiguration configuration, ILogger<GitHubAuthController> logger) : ControllerBase
    {
        [HttpGet("test")]
        public async Task<IActionResult> GitHubUserInfo()
        {
            await HttpContext.GetTokenAsync("acess_token");

            var externalId = HttpContext.User.Claims.FirstOrDefault(x=>x.Type == "externalId");
            var email = HttpContext.User.Claims.FirstOrDefault(x=>x.Type == "email");
            if (email is not null && externalId is not null)
                logger.LogInformation($"logged in [ ExternalId: {externalId.Value.AsSpan().ToString()}, "
                                                + $"Email: {email.Value.AsSpan().ToString()} ]");
            return Ok(
                HttpContext.User.Claims
                    .Select(c => new{c.Type, c.Value})
                    .ToList()
            );
        }

        [HttpGet("login")]
        public IActionResult GitHubLogin()
        {
            return Challenge
            (
                new AuthenticationProperties()
                {
                    RedirectUri = $"http://localhost:{configuration["BACKEND_HTTP_PORT"]}/api/auth/github/test",
                },
                authenticationSchemes:["GitHub"]
            );
        }

        [HttpGet("logout")]
        public async Task<IActionResult> Logout()
        {
            await HttpContext.SignOutAsync("GitHub");
            return Ok(new { message = "Logged out successfully" });
        }
    }
}
