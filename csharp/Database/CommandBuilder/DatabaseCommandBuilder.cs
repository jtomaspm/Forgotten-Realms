using System.Runtime.CompilerServices;
using System.Text;
using MySql.Data.MySqlClient;

namespace Database.CommandBuilder;

public static class DatabaseCommandBuilder
{
    public static MySqlCommand Select(this MySqlCommand cmd, IEnumerable<string> fields, string table, IEnumerable<string>? conditions, Dictionary<string, object>? parameters)
    {
        var conditionString = conditions is null ? "" : $"WHERE {ToConditionListString(conditions)}";
        var query = @$"SELECT {ToFieldnameListString(fields)} FROM {GetSafePropName(table)} {conditionString};";

        cmd.CommandText = query;
        if (parameters is not null)
        foreach (var parameter in parameters)
        {
            cmd.Parameters.AddWithValue(parameter.Key, parameter.Value);
        }
        return cmd;
    }
    public static MySqlCommand Insert(this MySqlCommand cmd, IEnumerable<string> fields, string table, IEnumerable<IEnumerable<string>> values, Dictionary<string, object>? parameters)
    {
        var query = @$"INSERT INTO {GetSafePropName(table)} {GetSafeConditionName(ToFieldnameListString(fields))} VALUES {table} {GetInsertValuesString(values)};";

        cmd.CommandText = query;
        if (parameters is not null)
        foreach (var parameter in parameters)
        {
            cmd.Parameters.AddWithValue(parameter.Key, parameter.Value);
        }
        return cmd;
    }
    private static string GetSafePropName (string name) 
        => $"`{name}`";
    private static string GetSafeConditionName (string name) 
        => $"({name})";
    private static string ToValueListString(IEnumerable<string> values) 
        => GetSafeConditionName(new StringBuilder().AppendJoin(", ", values).ToString());
    private static string ToFieldnameListString(IEnumerable<string> fields) 
        => new StringBuilder().AppendJoin(", ", fields.Select(GetSafePropName)).ToString();
    private static string ToConditionListString(IEnumerable<string> conditions) 
        => new StringBuilder().AppendJoin(" AND ", conditions.Select(GetSafeConditionName)).ToString();
    private static string GetInsertValuesString(IEnumerable<IEnumerable<string>> values)
        => ToFieldnameListString(values.Select(ToValueListString));
}
