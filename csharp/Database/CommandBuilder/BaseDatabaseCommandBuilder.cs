using System;
using System.Text;
using MySql.Data.MySqlClient;

namespace Database.CommandBuilder;

public abstract class BaseDatabaseCommandBuilder : IDatabaseCommandBuilder
{
    protected MySqlCommand _cmd;
    public BaseDatabaseCommandBuilder()
    {
        _cmd = new();
    }
    public BaseDatabaseCommandBuilder(MySqlCommand cmd)
    {
        _cmd = cmd;
    }
    public abstract MySqlCommand Build();
    protected static string SurroundWithCarrots (string str) 
        => $"`{str}`";
    protected static string SurroundWithParenthesis (string str) 
        => $"({str})";
    protected static string SeparateWithCommas(IEnumerable<string> strs) 
        => SurroundWithParenthesis(new StringBuilder().AppendJoin(", ", strs).ToString());
    protected static string SurroundWithCarrotsAndSeparateWithCommas(IEnumerable<string> strs) 
        => new StringBuilder().AppendJoin(", ", strs.Select(SurroundWithCarrots)).ToString();
    protected static string SeparateConditions(IEnumerable<string> conditions) 
        => new StringBuilder().AppendJoin(" AND ", conditions.Select(SurroundWithParenthesis)).ToString();
    protected static string GetValuesString(IEnumerable<IEnumerable<string>> values)
        => SurroundWithParenthesis(SeparateWithCommas(values.Select(value=>SurroundWithParenthesis(SeparateWithCommas(value)))));
}
