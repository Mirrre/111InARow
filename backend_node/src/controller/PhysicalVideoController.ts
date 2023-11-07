import { Body, Controller, Get, Post, Query } from '@snow';
import { PhysicalVideoDao } from '../dao/PhysicalVideoDao';
import { LikeVideoDao } from '../dao/LikeVideoDao';
import { CollectVideoDao } from '../dao/CollectVideoDao';

@Controller('/physical_video')
export default class PhysicalVideoController {
  private readonly physical_videoDao = new PhysicalVideoDao();

  @Post('/getPhysicalVideo')
  async getPhysicalVideo(@Body() body) {
    // return await this.physical_videoDao.getPhysicalVideo();
     // console.log(body);
     const result2 = await this.physical_videoDao.getPhysicalVideo();
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
        //  console.log(result2);
         
     }
     return result2;
  }

  @Post('/addPhysicalVideo')
  async addPhysicalVideo(@Body() body) {
    return await this.physical_videoDao.insert(body);
  }

  @Post('/findPhysicalVideoById')
  async findPhysicalVideo(@Query('id') id) {
    return await this.physical_videoDao.selectOne(id);
  }

  @Post('/findPhysicalVideoList')
  async findPhysicalVideoList(@Body() body) {
    const { page, size, ...queryParams } = body;
    return await this.physical_videoDao.selectPage(queryParams, { page, size }, { dirc: 'desc', field: 'ctime' });
  }

  @Post('/editPhysicalVideo')
  async editPhysicalVideo(@Body() body) {
    return await this.physical_videoDao.update(body.id, body);
  }

  @Post('/deletePhysicalVideo')
  async deletePhysicalVideo(@Query('id') id) {
    return await this.physical_videoDao.deleteById(id);
  }
}
