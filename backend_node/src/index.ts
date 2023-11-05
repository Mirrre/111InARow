import { Application, useMqtt } from '@snow';


// 设置 MQTT 连接选项，包括 Clean Session 标志
const mqttOptions = {
  clean: false, // 设置为 false 表示不清除会话状态
};

export const globEventEmitter = new (require('events'))();




@Application()
class SysApplication {
  before(app) {

  }
}

