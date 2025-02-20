using System;
using MySql.Data.MySqlClient;
using MySqlX.XDevAPI.Relational;

namespace Database.CommandBuilder;

public class SelectCommandBuilder : BaseDatabaseCommandBuilder
{
    private List<string> _fields = [];
    private string? _table;
    IEnumerable<string>? _conditions; 
    Dictionary<string, object>? _parameters;
    public SelectCommandBuilder():base(){}
    public SelectCommandBuilder(MySqlCommand cmd):base(cmd){}
    public SelectCommandBuilder Table(string table)
    {
        _table = table;
        return this;
    }
    public SelectCommandBuilder ClearFields() 
    {
        _fields=[];
        return this;
    }
    public SelectCommandBuilder AddField(string field)
    {
        _fields.Add(field);
        return this;
    }
    public SelectCommandBuilder AddFields(IEnumerable<string> fields)
    {
        _fields.AddRange(fields);
        return this;
    }
    public SelectCommandBuilder ClearConditions() 
    {
        _conditions=null;
        return this;
    }
    public SelectCommandBuilder AddConditions(IEnumerable<string> conditions)
    {
        _conditions ??= [.. conditions]; 
        return this;
    }
    public SelectCommandBuilder AddCondition(string condition)
    {
        _conditions ??= [condition]; 
        return this;
    }
    public SelectCommandBuilder SetParameters(Dictionary<string, object>? parameters)
    {
        _parameters = parameters;
        return this;
    }
    public override MySqlCommand Build()
    {
        if (_table is null) throw new ArgumentNullException("Trying to build Select Database Command without table set.");
        if (_fields.Count == 0) throw new ArgumentException("Trying to build Select Database Command without fields selt.");

        var conditionString = _conditions is null ? "" : $"WHERE {SeparateConditions(_conditions)}";
        var query = @$"SELECT {SurroundWithCarrotsAndSeparateWithCommas(_fields)} FROM {SurroundWithCarrots(_table)} {conditionString};";

        _cmd.CommandText = query;
        if (_parameters is not null)
        foreach (var parameter in _parameters)
        {
            _cmd.Parameters.AddWithValue(parameter.Key, parameter.Value);
        }
        return _cmd;
    }
}
