using System;
using MySql.Data.MySqlClient;

namespace Database.CommandBuilder;

public class InsertCommandBuilder : BaseDatabaseCommandBuilder
{
    private List<string> _fields = [];
    private string? _table;
    List<List<string>> _values = []; 
    Dictionary<string, object>? _parameters;
    public InsertCommandBuilder():base(){}
    public InsertCommandBuilder(MySqlCommand cmd):base(cmd){}
    public InsertCommandBuilder Table(string table)
    {
        _table = table;
        return this;
    }
    public InsertCommandBuilder ClearFields() 
    {
        _fields=[];
        return this;
    }
    public InsertCommandBuilder AddField(string field)
    {
        _fields.Add(field);
        return this;
    }
    public InsertCommandBuilder AddFields(IEnumerable<string> fields)
    {
        _fields.AddRange(fields);
        return this;
    }
    public InsertCommandBuilder ClearValues() 
    {
        _values=[];
        return this;
    }
    public InsertCommandBuilder AddValues(IEnumerable<IEnumerable<string>> values)
    {
        _values.AddRange(values.Select(x => x.ToList()));
        return this;
    }
    public InsertCommandBuilder AddValueSet(IEnumerable<string> value)
    {
        _values.Add([.. value]); 
        return this;
    }
    public InsertCommandBuilder SetParameters(Dictionary<string, object>? parameters)
    {
        _parameters = parameters;
        return this;
    }
    public override MySqlCommand Build()
    {
        if (_table is null) throw new ArgumentNullException("Trying to build Insert Database Command without table set.");
        if (_fields.Count == 0) throw new ArgumentException("Trying to build Insert Database Command without fields selt.");
        if (_values.Count == 0) throw new ArgumentException("Trying to build Insert Database Command without values selt.");

        var query = @$"INSERT INTO {SurroundWithCarrots(_table)} {SurroundWithParenthesis(SurroundWithCarrotsAndSeparateWithCommas(_fields))} VALUES {GetValuesString(_values)};";

        _cmd.CommandText = query;
        if (_parameters is not null)
        foreach (var parameter in _parameters)
        {
            _cmd.Parameters.AddWithValue(parameter.Key, parameter.Value);
        }
        return _cmd;
    }

}
