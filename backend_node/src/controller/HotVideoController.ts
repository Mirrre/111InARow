import { Body, Controller, Get, Post, Query } from '@snow';
import { HotVideoDao } from '../dao/HotVideoDao';
import { LikeVideoDao } from '../dao/LikeVideoDao';
import { CollectVideoDao } from '../dao/CollectVideoDao';

@Controller('/hot_video')
export default class HotVideoController {
  private readonly hot_videoDao = new HotVideoDao();

  @Post('/getHotVideo')
  async getHotVideo(@Body() body) {
    // console.log(body);
    const result2 = await this.hot_videoDao.getHotVideo();
    // console.log(result2);
    
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

  @Post('/addHotVideo')
  async addHotVideo(@Body() body) {
    return await this.hot_videoDao.insert(body);
  }

  @Post('/findHotVideoById')
  async findHotVideo(@Query('id') id) {
    return await this.hot_videoDao.selectOne(id);
  }

  @Post('/findHotVideoList')
  async findHotVideoList(@Body() body) {
    const { page, size, ...queryParams } = body;
    return await this.hot_videoDao.selectPage(queryParams, { page, size }, { dirc: 'desc', field: 'ctime' });
  }

  @Post('/editHotVideo')
  async editHotVideo(@Body() body) {
    return await this.hot_videoDao.update(body.id, body);
  }

  @Post('/deleteHotVideo')
  async deleteHotVideo(@Query('id') id) {
    return await this.hot_videoDao.deleteById(id);
  }
}
