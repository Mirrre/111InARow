import { Body, Controller, Get, Post, Query } from '@snow';
import { UsersDao } from '../dao/UsersDao';

@Controller('/users')
export default class UsersController {
  private readonly usersDao = new UsersDao();

  @Post('/login')
  async findUserByUsername(@Body() body) {
    console.log(body);
    return await this.usersDao.findUserByUsername(body.username,body.password);
  }

  @Post('/register')
  async addUsers(@Body() body) {
    console.log(body);
    await this.usersDao.insert(body);
    const result2 = await this.usersDao.findUserByUsername(body.username,body.password);
    return result2;
  }

  @Post('/findUsersById')
  async findUsers(@Query('id') id) {
    return await this.usersDao.selectOne(id);
  }

  @Post('/findUsersList')
  async findUsersList(@Body() body) {
    const { page, size, ...queryParams } = body;
    return await this.usersDao.selectPage(queryParams, { page, size }, { dirc: 'desc', field: 'ctime' });
  }

  @Post('/editUsers')
  async editUsers(@Body() body) {
    return await this.usersDao.update(body.id, body);
  }

  @Post('/deleteUsers')
  async deleteUsers(@Query('id') id) {
    return await this.usersDao.deleteById(id);
  }
}
