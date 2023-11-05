import { Body, Controller, Get, Post, Query } from '@snow';
import { GameVideoDao } from '../dao/GameVideoDao';
import { LikeVideoDao } from '../dao/LikeVideoDao';
import { CollectVideoDao } from '../dao/CollectVideoDao';

@Controller('/game_video')
export default class GameVideoController {
  private readonly game_videoDao = new GameVideoDao();

  @Post('/getGameVideo')
  async getGameVideo(@Body() body) {
    // return await this.game_videoDao.getGameVideo();
    // console.log(body);
    const result2 = await this.game_videoDao.getGameVideo();
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

  @Post('/addGameVideo')
  async addGameVideo(@Body() body) {
    return await this.game_videoDao.insert(body);
  }

  @Post('/findGameVideoById')
  async findGameVideo(@Query('id') id) {
    return await this.game_videoDao.selectOne(id);
  }

  @Post('/findGameVideoList')
  async findGameVideoList(@Body() body) {
    const { page, size, ...queryParams } = body;
    return await this.game_videoDao.selectPage(queryParams, { page, size }, { dirc: 'desc', field: 'ctime' });
  }

  @Post('/editGameVideo')
  async editGameVideo(@Body() body) {
    return await this.game_videoDao.update(body.id, body);
  }

  @Post('/deleteGameVideo')
  async deleteGameVideo(@Query('id') id) {
    return await this.game_videoDao.deleteById(id);
  }
}
