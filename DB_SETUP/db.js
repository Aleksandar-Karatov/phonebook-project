// mongoose.connect(process.env.DB_CONNECTION, { useNewUrlParser: true, useUnifiedTopology: true });

const express = require("express");
const mongoose = require("mongoose");
const app = express();

const uri = "mongodb+srv://Team8:1234566678@cluster0.ssiamrl.mongodb.net/?retryWrites=true&w=majority";

async function connect() {
  try {
    await mongoose.connect(uri);
    console.log("Connected to MongoDB");
  } catch (error) {
    console.error(error);
  }
}

connect();
console.log(uri);

app.listen(3000, () => {
  console.log("Server started on port 3000");
});

const phoneBookSchema = new mongoose.Schema({
  id: {
    type: String,
    required: true,
    unique: true
  },
  name: String,
  phoneNumbers: [String],
  notes: String,
  email: String
});

const PhoneBook = mongoose.model('PhoneBook', phoneBookSchema);

async function insertPhoneBookEntry(entry) {
  try {
    const newEntry = new PhoneBook({
      id: entry.id,
      name: entry.name,
      phoneNumbers: entry.phoneNumbers,
      notes: entry.notes,
      email: entry.email
    });

    const savedEntry = await newEntry.save();
    console.log('Inserted new phone book entry:', savedEntry);
  } catch (err) {
    console.log(err);
  }
}

function deletePhoneBookEntry(id) {
  PhoneBook.findOneAndDelete({ id: id }, function (err, entry) {
    if (err) {
      console.log(err);
    } else if (!entry) {
      console.log('No phone book entry found with ID:', id);
    } else {
      console.log('Deleted phone book entry:', entry);
    }
  });
}

async function findPhoneBookEntry(name) {
  try {
    const entries = await PhoneBook.find({ name: name });
    console.log('Found phone book entries:', entries);
  } catch (err) {
    console.log(err);
  }
}

async function updatePhoneBookEntry(name, changes) {
  try {
    const result = await PhoneBook.updateMany({ name: name }, changes);
    console.log('Updated phone book entries:', result);
  } catch (err) {
    console.log(err);
  }
}

const boyanna = {
  id: '1',
  name: 'Boyanna',
  phoneNumbers: ['123-456-7890'],
  notes: 'Student',
  email: 'boyanna@gmail.com'
};

insertPhoneBookEntry(boyanna);

const yana = {
  id: '2',
  name: 'Yana',
  phoneNumbers: ['123-456-7888', '123-456-7788'],
  notes: 'Student',
  email: 'yana@gmail.com'
};

insertPhoneBookEntry(yana);

const aleks = {
  id: '3',
  name: 'Aleks K',
  phoneNumbers: ['123-456-7800'],
  notes: 'Student',
  email: 'aleks@gmail.com'
};

insertPhoneBookEntry(aleks);

findPhoneBookEntry('Boyanna');

updatePhoneBookEntry('Yana', { notes: 'Lives in Sofia' });
