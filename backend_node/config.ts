import { ConfigOptions } from './snow/configuration';

const config: ConfigOptions = {
  port: 8901,
  rootPath: '/api',
  database: {
    host: 'rm-cn-9lb3g9sil00042do.rwlb.rds.aliyuncs.com',
    username: 'lips',
    port: 2396,
    password: 'yhy@PS4611',
    database: '111inarow', // 你的数据库名称
  },
};

// const config: ConfigOptions = {
//   port: 8901,
//   rootPath: '/api',
//   database: {
//   host: 'rm-cn-9lb3g9sil00042do.rwlb.rds.aliyuncs.com',
//   username: 'lips',
//   port: 2396,
//   password: 'yhy@PS4611',
//   database: '111inarow', // 你的数据库名称
//   },
// };
export default config;
