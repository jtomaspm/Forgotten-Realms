using MySql.Data.MySqlClient;

namespace Database.CommandBuilder;

public class DeleteCommandBuilder : BaseDatabaseCommandBuilder
{
    private string? _table;
    private IEnumerable<string>? _conditions;
    private Dictionary<string, object>? _parameters;

    public DeleteCommandBuilder() : base() { }
    public DeleteCommandBuilder(MySqlCommand cmd) : base(cmd) { }

    public DeleteCommandBuilder Table(string table)
    {
        _table = table;
        return this;
    }

    public DeleteCommandBuilder ClearConditions()
    {
        _conditions = null;
        return this;
    }

    public DeleteCommandBuilder AddConditions(IEnumerable<string> conditions)
    {
        _conditions ??= [.. conditions];
        return this;
    }

    public DeleteCommandBuilder AddCondition(string condition)
    {
        _conditions ??= [condition];
        return this;
    }

    public DeleteCommandBuilder SetParameters(Dictionary<string, object>? parameters)
    {
        _parameters = parameters;
        return this;
    }

    public override MySqlCommand Build()
    {
        if (_table is null) throw new ArgumentNullException("Trying to build Delete Database Command without table set.");
        if (_conditions is null || !_conditions.Any()) throw new ArgumentException("Trying to build Delete Database Command without conditions set.");

        var conditionString = $"WHERE {SeparateConditions(_conditions)}";
        var query = $"DELETE FROM {SurroundWithCarrots(_table)} {conditionString};";

        _cmd.CommandText = query;
        if (_parameters is not null)
        {
            foreach (var parameter in _parameters)
            {
                _cmd.Parameters.AddWithValue(parameter.Key, parameter.Value);
            }
        }

        return _cmd;
    }
}