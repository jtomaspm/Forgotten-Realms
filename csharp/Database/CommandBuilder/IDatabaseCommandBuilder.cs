using MySql.Data.MySqlClient;

namespace Database.CommandBuilder;

public interface IDatabaseCommandBuilder
{
    public MySqlCommand Build();
}
