using Microsoft.Extensions.Configuration;

namespace ApiUtils.Configuration;

public static class IConfigurationExtetions
{
    public static void ValidateEnv(this IConfiguration configuration, string envVarsList)
    {
        List<string> envVars = configuration.GetRequiredSection(envVarsList).AsEnumerable().Where(x=>x.Value is not null).Select(x=>x.Value).ToList()!;
        
        List<string> invalidFields = new();
        foreach (var field in envVars) 
            if (configuration[field] is null)
                invalidFields.Add(field);

        if (invalidFields.Count != 0)
        {
            var error = String.Empty;
            foreach (var field in invalidFields)
                error += $"{field} variable not set in .env file.\n";
            throw new ArgumentNullException(error);
        }
    }
}
