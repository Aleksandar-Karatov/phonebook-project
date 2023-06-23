import { createUser } from "./CRUD";

createUser(mario, marioImage)
    .then((createdUser) => {
        console.log("User created:", createdUser);
    })
    .catch((error) => {
        console.error("Error creating user:", error.message);
    });

getUserById('646b74ec5dd6e98d5e632c85')
    .then((user) => {
        console.log("User found:", user);
    })
    .catch((error) => {
        console.error("Error getting user:", error.message);
    });

updateUser("646b74ec5dd6e98d5e632c85", { name: "Mihail", phoneNumbers: "0877743210" })
    .then((userData) => {
        console.log("User updated:", userData);
    })
    .catch((error) => {
        console.error("Error updating user:", error.message);
    });

deleteUser("64624140630984f8cbdd2831")
    .then((deletedUser) => {
        console.log("User deleted:", deletedUser);
    })
    .catch((error) => {
        console.error("Error deleting user:", error.message);
    });

module.exports = {
    createUser
};