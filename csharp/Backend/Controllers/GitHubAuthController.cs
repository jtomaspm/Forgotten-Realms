using Microsoft.AspNetCore.Authentication;
using Microsoft.AspNetCore.Authentication.Cookies;
using Microsoft.AspNetCore.Mvc;
using MySql.Data.MySqlClient;
using System.Security.Claims;

namespace Backend.Controllers
{
    [Route("api/auth/github")]
    [ApiController]
    public class GitHubAuthController(IConfiguration configuration) : ControllerBase
    {
        [HttpGet("test")]
        public async Task<IActionResult> GitHubUserInfo()
        {
            await HttpContext.GetTokenAsync("acess_token");
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
