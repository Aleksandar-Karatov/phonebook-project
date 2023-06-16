const { default: mongoose } = require("mongoose");
const { User, Image } = require('./schemas');

async function createUser(userData, image) {
  try {
    const newImage = new Image({
      name: image.name,
      data: image.data,
      contentType: image.contentType
    });

    const savedImage = await newImage.save();

    const newUser = new User({
      id: new mongoose.Types.ObjectId(),
      name: userData.name,
      phoneNumbers: userData.phoneNumbers,
      notes: userData.notes,
      email: userData.email,
      image_id: savedImage._id
    });

    const savedUser = await newUser.save();
    return savedUser;
  } catch (error) {
    throw new Error("Failed to create user: " + error.message);
  }
}

async function getUserById(userId) {
  try {
    const user = await User.findById(userId);
    if (!user) {
      throw new Error("User not found");
    }
    return user;
  } catch (error) {
    throw new Error("Failed to get user: " + error.message);
  }
}

async function updateUser(userId, userData) {
  try {
    const updatedUser = await User.findByIdAndUpdate(userId, userData, {
      new: true,
    });
    if (!updatedUser) {
      throw new Error("User not found");
    }
    return updatedUser;
  } catch (error) {
    throw new Error("Failed to update user: " + error.message);
  }
}

async function deleteUser(userId) {
  try {
    const deletedUser = await User.findByIdAndRemove(userId);
    if (!deletedUser) {
      throw new Error("User not found");
    }
    return deletedUser;
  } catch (error) {
    throw new Error("Failed to delete user: " + error.message);
  }
}

module.exports = {
  createUser,
  getUserById,
  updateUser,
  deleteUser,
};