const app = express();

createUser(mario, marioImage)
    .then((createdUser) => {
        console.log("User created:", createdUser);
    })
    .catch((error) => {
        console.error("Error creating user:", error.message);
    });

getUserById("6485ac44117343ba41b74c64");
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

app.get('/users/:id/image', async (req, res) => {
    try {
        const userId = req.params.id;

        const user = await User.findById(userId).exec();

        if (!user) {
            return res.status(404).send('User not found');
        }

        if (!user.image_id) {
            return res.status(404).send('Image not found for the user');
        }

        const image = await Image.findById(user.image_id).exec();

        if (!image) {
            return res.status(404).send('Image not found');
        }

        res.set('Content-Type', image.contentType);

        res.send(image.data);
    } catch (error) {
        console.error('Error retrieving image:', error);
        res.status(500).send('Internal Server Error');
    }
});