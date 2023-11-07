import { Body, Controller, Get, Post, Query } from '@snow';
import { CollectVideoDao } from '../dao/CollectVideoDao';
import { AllVideoDao } from '../dao/AllVideoDao';

@Controller('/collect_video')
export default class CollectVideoController {
  private readonly collect_videoDao = new CollectVideoDao();

  @Post('/deleteCollectVideoByLimit')
  async deleteCollectVideoByLimit(@Body() body) {
    // console.log(body);
    const sql2 = `UPDATE all_video SET collect_num = collect_num - 1 WHERE id = ?`
    const id = body.video_id
    await AllVideoDao.query(sql2,[id])
    const sql = `delete  from collect_video where user_id = ? and video_id = ?`
    return await CollectVideoDao.query(sql,[body.user_id,body.video_id])
  }

  @Post('/getCollectVideoByLimit')
  async getCollectVideoByLimit(@Body() body) {
    // console.log(body);
    
    const sql = `select * from collect_video where user_id = ? and video_id = ?`
    return await CollectVideoDao.query(sql,[body.user_id,body.video_id])
  }

  @Post('/addCollectVideo')
  async addCollectVideo(@Body() body) {
    const sql = `UPDATE all_video SET collect_num = collect_num + 1 WHERE id = ?`
    const id = body.video_id
    await AllVideoDao.query(sql,[id])
    return await this.collect_videoDao.insert(body);
  }

  @Post('/findCollectVideoById')
  async findCollectVideo(@Query('id') id) {
    return await this.collect_videoDao.selectOne(id);
  }

  @Post('/findCollectVideoList')
  async findCollectVideoList(@Body() body) {
    const { page, size, ...queryParams } = body;
    return await this.collect_videoDao.selectPage(queryParams, { page, size }, { dirc: 'desc', field: 'ctime' });
  }

  @Post('/editCollectVideo')
  async editCollectVideo(@Body() body) {
    return await this.collect_videoDao.update(body.id, body);
  }

  @Post('/deleteCollectVideo')
  async deleteCollectVideo(@Query('id') id) {
    return await this.collect_videoDao.deleteById(id);
  }
}
