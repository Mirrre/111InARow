import { BaseDao, Dao, Sql } from '@snow';

@Dao('users')
export class UsersDao extends BaseDao {
    @Sql(`select * from users where username = ? and password = ?`)
  async findUserByUsername(username,password): Promise<any> {}
}
