import { BaseDao, Dao, Sql } from '@snow';

@Dao('game_video')
export class GameVideoDao extends BaseDao {
    // @Sql(`select * from game_video`)
    // @Sql(`SELECT all_video.* FROM all_video LEFT JOIN game_video ON all_video.id = game_video.all_id;`)
    @Sql(`SELECT * FROM game_video JOIN videos ON videos.id = game_video.all_id`)
    async getGameVideo(): Promise<any> {}
}
