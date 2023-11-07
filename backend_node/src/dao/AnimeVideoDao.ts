import { BaseDao, Dao, Sql } from '@snow';

@Dao('anime_video')
export class AnimeVideoDao extends BaseDao {
    // @Sql(`select * from anime_video`)
    // @Sql(`SELECT all_video.* FROM all_video LEFT JOIN anime_video ON all_video.id = anime_video.all_id;`)
    @Sql(`SELECT * FROM anime_video JOIN videos ON videos.id = anime_video.all_id`)
    async getAnimeVideo(): Promise<any> {}
}
