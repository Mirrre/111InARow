import { Body, Controller, Get, Post, Query } from '@snow';
import { AllVideoDao } from '../dao/AllVideoDao';

@Controller('/all_video')
export default class AllVideoController {
  private readonly all_videoDao = new AllVideoDao();

  @Post('/findAllVideo')
  async findAllVideo(@Body() body) {
    return await this.all_videoDao.findAllVideo();
  }

  @Post('/addAllVideo')
  async addAllVideo(@Body() body) {
    return await this.all_videoDao.insert(body);
  }

  @Post('/findAllVideoById')
  async findAllVideoById(@Query('id') id) {
    return await this.all_videoDao.selectOne(id);
  }

  @Post('/findAllVideoList')
  async findAllVideoList(@Body() body) {
    const { page, size, ...queryParams } = body;
    return await this.all_videoDao.selectPage(queryParams, { page, size }, { dirc: 'desc', field: 'ctime' });
  }

  @Post('/editAllVideo')
  async editAllVideo(@Body() body) {
    return await this.all_videoDao.update(body.id, body);
  }

  @Post('/deleteAllVideo')
  async deleteAllVideo(@Query('id') id) {
    return await this.all_videoDao.deleteById(id);
  }
}
