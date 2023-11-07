import { BaseDao, Dao, Sql } from '@snow';

@Dao('all_video')
export class AllVideoDao extends BaseDao {
    @Sql(`select * from videos`)
  async findAllVideo(): Promise<any> {}
}
