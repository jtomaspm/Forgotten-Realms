using System.Text;
using MySql.Data.MySqlClient;

namespace Database.CommandBuilder;

public static class DatabaseCommandBuilder
{
    private static string GetSafePropName (string name) => $"`{name}`";
    private static string GetSafeConditionName (string name) => $"({name})";
    private static string ToFieldListString(IEnumerable<string> names) 
    {
        var sb = new StringBuilder();
        sb.AppendJoin(", ", names.Select(x=>GetSafePropName(x)));
        return sb.ToString();
    }
    private static string ToConditionListString(IEnumerable<string> names) 
    {
        var sb = new StringBuilder();
        sb.AppendJoin(" AND ", names.Select(x=>GetSafeConditionName(x)));
        return sb.ToString();
    }
    public static MySqlCommand Select(this MySqlCommand cmd, IEnumerable<string> fields, string table, IEnumerable<string>? conditions, Dictionary<string, object>? parameters)
    {
        var conditionString = conditions is null ? "" : $"WHERE {ToConditionListString(conditions)}";
        var query = @$"SELECT {ToFieldListString(fields)} FROM {table} {conditionString};";

        cmd.CommandText = query;
        if (parameters is not null)
        foreach (var parameter in parameters)
        {
            cmd.Parameters.AddWithValue(parameter.Key, parameter.Value);
        }
        return cmd;
    }
}
