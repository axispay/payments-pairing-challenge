const express=require('express');function createApp(router){const app=express();app.use(express.json());app.use(router);return app}module.exports={createApp};
