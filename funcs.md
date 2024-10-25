This is a list of functions that are needed and will be implemented on the server side.
Not in order but a collection.
For reference (not documentation).

- Will have to take an image from frontend through the request and perform some verifications and send it to the appropriate input receiver.
- Will check for the output in database and retrieve it and send it to the frontend when requested.
- Will take the necessary 3d env data and send it to the frontend when requested.
- Will take the updated 3d env data and sends it to the function related and takes the output and passes it onto the rendering engine.
- Will take the 3d env data and call the rendering engine on the data located. Once the rendered image is given, some verifications will be done and the image location is stored in database.
- Will keep checking if the output is done yet. Once done, the image is sent to the frontend.
- Will mark the task as done and clears it in the database. Once done, all the files related will be removed.

Endpoints/HandlerFuncs:

- One for the status check.
- One for the input.
- One for the 3d env to frontend.
- One to get the updated 3d env.
- One for statuses and messages to the user.
