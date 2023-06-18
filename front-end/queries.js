export const fetchContacts = `
query {
    contacts() {
      name, id
    }
  }`
;

export var contact = `
query contact($name: String, $phone: String, $email: String, $id: String){
    contact(name: $name, phone: $phone, email: $email, id: $id){
      id
      name
      phones
      notes
      email
      imageID
    }
  }`
;

export var search = `
query search{
    search(filter: $filter){
      name
      id
    }
  }
`;

 export var image = `
query image{
    image(imageID:"$imageID"){
      imageId
      name
      imageData
      contentType
    }
  }
`;

export var createContact = `
mutation create{
    createContact(name: $name, phones: $phones, notes: $notes, email: $email){
      id
      name
      notes
      phones
      email
      imageID
    }
  }
`;

export var mergeContacts = `
mutation merge{
    mergeContacts(id: $id){
      name
      id
      email
      phones
      imageID
      notes
    }
  }
`;

export var updateContact = `
mutation create{
    updateContact(id: $id , name: $name, phones: $phones, notes: $notes, email: $email){
      id
      name
      notes
      phones
      email
      imageID
    }
  }
`;

export var deleteContact = `
    mutation{
      deleteContact(id: $id)
    }
`

export var updateContactPicture = `
    query{
        updateContactPicture(userId: $userID, imageId: $imageID, name: $name, imageData: $imageData, contentType: $contentType){
            imageId
            name
            imageData
            contentType
        }
    }
`