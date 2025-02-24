using Dapper;
using Database.Application.Models;
using System.Data;

namespace Database.Application.TypeHandlers;

public class RoleTypeHandler : SqlMapper.TypeHandler<Role>
{
    public override Role Parse(object value)
    {
        return Role.FromNameString(value.ToString()!);
    }

    public override void SetValue(IDbDataParameter parameter, Role? value)
    {
        if (value is null) 
        {
            parameter.Value = null;
            return;
        }
        parameter.Value = value.Name.ToString();
    }
}