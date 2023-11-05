import { Body, Controller, Get, Post, Query } from '@snow';
import { LikeVideoDao } from '../dao/LikeVideoDao';
import { AllVideoDao } from '../dao/AllVideoDao';

@Controller('/like_video')
export default class LikeVideoController {
  private readonly like_videoDao = new LikeVideoDao();

  @Post('/getLikeVideoByUser')
  async getLikeVideoByUser(@Body() body) {
    // console.log(body);
    const sql = `SELECT all_video.* FROM all_video LEFT JOIN like_video ON all_video.id = like_video.video_id where like_video.user_id = ?`
    const data = await LikeVideoDao.query(sql,[body.user_id])
    // console.log(data);
    
    return await LikeVideoDao.query(sql,[body.user_id])
  }

  @Post('/deleteLikeVideoByLimit')
  async deleteLikeVideoByLimit(@Body() body) {
    // console.log(body);
    const sql2 = `UPDATE all_video SET like_num = like_num - 1 WHERE id = ?`
    const id = body.video_id
    await AllVideoDao.query(sql2,[id])
    const sql = `delete  from like_video where user_id = ? and video_id = ?`
    return await LikeVideoDao.query(sql,[body.user_id,body.video_id])
  }

  @Post('/getLikeVideoByLimit')
  async getLikeVideoByLimit(@Body() body) {
    // console.log(body);
    
    const sql = `select * from like_video where user_id = ? and video_id = ?`
    return await LikeVideoDao.query(sql,[body.user_id,body.video_id])
  }

  @Post('/addLikeVideo')
  async addLikeVideo(@Body() body) {
    // console.log(body);
    const sql = `UPDATE all_video SET like_num = like_num + 1 WHERE id = ?`
    const id = body.video_id
    await AllVideoDao.query(sql,[id])
    return await this.like_videoDao.insert(body);
  }

  @Post('/findLikeVideoById')
  async findLikeVideo(@Query('id') id) {
    return await this.like_videoDao.selectOne(id);
  }

  @Post('/findLikeVideoList')
  async findLikeVideoList(@Body() body) {
    const { page, size, ...queryParams } = body;
    return await this.like_videoDao.selectPage(queryParams, { page, size }, { dirc: 'desc', field: 'ctime' });
  }

  @Post('/editLikeVideo')
  async editLikeVideo(@Body() body) {
    return await this.like_videoDao.update(body.id, body);
  }

  @Post('/deleteLikeVideo')
  async deleteLikeVideo(@Query('id') id) {
    return await this.like_videoDao.deleteById(id);
  }
}
