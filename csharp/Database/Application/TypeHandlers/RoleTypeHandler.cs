using Dapper;
using Database.Application.Models;
using System.Data;

namespace Database.Application.TypeHandlers;

public class RoleTypeHandler : SqlMapper.TypeHandler<Role>
{
    public override Role Parse(object value) =>
        Role.FromNameString(value.ToString()!);

    public override void SetValue(IDbDataParameter parameter, Role? value)
    {
        parameter.Value = value?.Name.ToString();
    }
}