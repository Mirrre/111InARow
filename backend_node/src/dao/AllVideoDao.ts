import { BaseDao, Dao, Sql } from '@snow';

@Dao('all_video')
export class AllVideoDao extends BaseDao {
    @Sql(`select * from all_video`)
  async findAllVideo(): Promise<any> {}
}
