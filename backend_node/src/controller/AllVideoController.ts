import { Body, Controller, Get, Post, Query } from '@snow';
import { AllVideoDao } from '../dao/AllVideoDao';
import { LikeVideoDao } from '../dao/LikeVideoDao';
import { CollectVideoDao } from '../dao/CollectVideoDao';

@Controller('/all_video')
export default class AllVideoController {
  private readonly all_videoDao = new AllVideoDao();

  @Post('/findUploadVideoById')
  async findUploadVideoById(@Body() body) {
    const sql = `select * from videos where user_id = ?`
    return await AllVideoDao.query(sql, [body.user_id]);
  }

  @Post('/findAllVideoByTitle')
  async findAllVideoByTitle(@Body() body) {
    // console.log(body);
    const title = '%' + body.title + '%';
    const sql2 = `select * from videos where title like ? `;
    const result2 = await AllVideoDao.query(sql2, [title]);
    if ('user_id' in body) {
      const sql = `select * from like_video where user_id = ?`;
      const result1 = await LikeVideoDao.query(sql, [body.user_id]);
      const sql3 = `select * from collect_video where user_id = ?`;
      const result3 = await CollectVideoDao.query(sql3, [body.user_id]);
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

  @Post('/findAllVideo')
  async findAllVideo(@Body() body) {
    // const result2 = await this.all_videoDao.findAllVideo();
    // console.log(result2);

    // return await this.all_videoDao.findAllVideo();
    const result2 = await this.all_videoDao.findAllVideo();
    if ('has' in body) {
      const sql = `select * from like_video where user_id = ?`;
      const result1 = await LikeVideoDao.query(sql, [body.user_id]);
      const sql3 = `select * from collect_video where user_id = ?`;
      const result3 = await CollectVideoDao.query(sql3, [body.user_id]);
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
