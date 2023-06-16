const mongoose = require("mongoose");

const UserSchema = new mongoose.Schema({
    id: {
        type: mongoose.Schema.Types.ObjectId,
        required: true,
        unique: true
    },
    name: {
        type: mongoose.Schema.Types.String,
        required: true
    },
    phoneNumbers: {
        type: mongoose.Schema.Types.Array,
        required: true,
    },
    notes: {
        type: mongoose.Schema.Types.String,
        required: true,
    },
    email: {
        type: mongoose.Schema.Types.String,
        required: true,
        match: [/^[a-zA-Z0-9.!#$%&’*+/=?^_`{|}~-]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*$/, 'Invalid email']
    },
    image_id: {
        type: mongoose.Schema.Types.ObjectId,
        ref: 'Image'
    }
});

const ImageSchema = new mongoose.Schema({
    name: {
      type: String,
      required: true
    },
    data: {
      type: Buffer,
      required: true
    },
    contentType: {
      type: String,
      required: true
    }
});

const Image = mongoose.model('Image', ImageSchema);
const User = mongoose.model('User', UserSchema);

module.exports = { Image, User };