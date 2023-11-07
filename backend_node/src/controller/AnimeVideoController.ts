import { Body, Controller, Get, Post, Query } from '@snow';
import { AnimeVideoDao } from '../dao/AnimeVideoDao';
import { LikeVideoDao } from '../dao/LikeVideoDao';
import { CollectVideoDao } from '../dao/CollectVideoDao';

@Controller('/anime_video')
export default class AnimeVideoController {
  private readonly anime_videoDao = new AnimeVideoDao();

  @Post('/getAnimeVideo')
  async getAnimeVideo(@Body() body) {
    // return await this.anime_videoDao.getAnimeVideo();
        // console.log(body);
        const result2 = await this.anime_videoDao.getAnimeVideo();
        if('has' in body){
            const sql = `select * from like_video where user_id = ?`
            const  result1 =  await LikeVideoDao.query(sql,[body.user_id])
            const sql3 = `select * from collect_video where user_id = ?`
            const  result3 =  await CollectVideoDao.query(sql3,[body.user_id])
            // console.log(result1);
            // console.log(result3);
            for (const item1 of result1) {
              for (const item2 of result2) {
                if (item1.video_id === item2.id) {
                  item2.is_follow = '1';
                }
              }
            }
            for (const item3 of result3) {
              for (const item2 of result2) {
                if (item3.video_id === item2.id) {
                  item2.is_follow2 = '1';
                }
              }
            }
            // console.log(result2);
            
        }
        return result2;
  }

  @Post('/addAnimeVideo')
  async addAnimeVideo(@Body() body) {
    return await this.anime_videoDao.insert(body);
  }

  @Post('/findAnimeVideoById')
  async findAnimeVideo(@Query('id') id) {
    return await this.anime_videoDao.selectOne(id);
  }

  @Post('/findAnimeVideoList')
  async findAnimeVideoList(@Body() body) {
    const { page, size, ...queryParams } = body;
    return await this.anime_videoDao.selectPage(queryParams, { page, size }, { dirc: 'desc', field: 'ctime' });
  }

  @Post('/editAnimeVideo')
  async editAnimeVideo(@Body() body) {
    return await this.anime_videoDao.update(body.id, body);
  }

  @Post('/deleteAnimeVideo')
  async deleteAnimeVideo(@Query('id') id) {
    return await this.anime_videoDao.deleteById(id);
  }
}
