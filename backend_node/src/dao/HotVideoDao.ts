import { BaseDao, Dao, Sql } from '@snow';

@Dao('hot_video')
export class HotVideoDao extends BaseDao {
    @Sql(`SELECT all_video.* FROM all_video LEFT JOIN hot_video ON all_video.id = hot_video.all_id;`)
    async getHotVideo(): Promise<any> {}
}
