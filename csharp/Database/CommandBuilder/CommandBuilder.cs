using System.Text;
using MySql.Data.MySqlClient;

namespace Database.CommandBuilder;

public class CommandBuilder
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
    public static MySqlCommand Select(IEnumerable<string> fields, string table, IEnumerable<string>? conditions, Dictionary<string, object>? parameters)
    {
        var cmd = new MySqlCommand();

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
