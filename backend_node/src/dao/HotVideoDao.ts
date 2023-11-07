import { BaseDao, Dao, Sql } from '@snow';

@Dao('hot_video')
export class HotVideoDao extends BaseDao {
    @Sql(`SELECT * FROM hot_video JOIN videos ON videos.id = hot_video.all_id`)
    async getHotVideo(): Promise<any> {}
}
