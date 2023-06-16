const express = require("express");
const mongoose = require("mongoose");
const { User, Image } = require('./schemas');
const app = express();
const dotenv = require('dotenv');
const { createUser, getUserById, updateUser, deleteUser } = require("./CRUD");
const { boyanna, yana, aleks, dobrin, nevena, yavor, ivo, gabriela,
        konstantin, mariya, kristiyan, milena, aleksandra, martin, mario, 
        boyannaImage, yanaImage, aleksImage, dobrinImage, nevenaImage,
        yavorImage, ivoImage, gabrielaImage, konstantinImage, mariyaImage,
        kristiyanImage, milenaImage, aleksndraImage, martinImage, marioImage} = require("./userData");


async function connect() {
  try {
    const result = dotenv.config();
    if (result.error) {
      throw result.error;
    }

    await mongoose.connect(process.env.DB_CONNECTION, {
      useNewUrlParser: true,
      useUnifiedTopology: true,
    });
    console.log('Connected to MongoDB');
    startServer();
  } catch (error) {
    console.error(error);
  }
}
  
connect();

function startServer() {
    app.listen(3000, () => {
        console.log("Server started on port 3000");
    });
}