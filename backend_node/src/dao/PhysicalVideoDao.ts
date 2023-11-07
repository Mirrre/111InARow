import { BaseDao, Dao, Sql } from '@snow';

@Dao('physical_video')
export class PhysicalVideoDao extends BaseDao {
    // @Sql(`select * from physical_video`)
    @Sql(`SELECT all_video.* FROM all_video LEFT JOIN physical_video ON all_video.id = physical_video.all_id;`)
    async getPhysicalVideo(): Promise<any> {}
}
