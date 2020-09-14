<!DOCTYPE html>
<html lang="en">
  <head>
    <title>File Upload</title>
  </head>
  <body>
    <form
      enctype="multipart/form-data"
      action="http://localhost:9000/recipe/stats" method="post"
    >
       PostCode&nbsp;<input type="text" id="postcode" name="postCode"><br>
       RecipeName&nbsp;<input type="text" id="recipeName" name="recipeName"><br>
       DeliveryStartTime&nbsp;<input type="text" id="deliveryStartTime" name="deliveryStartTime"><span>Ex. 01PM, 11AM</span><br>
       DeliveryEndTime&nbsp;<input type="text" id="deliveryEndTime" name="deliveryEndTime"><span>Ex. 01PM, 11AM</span><br>
       Upload File &nbsp;<input type="file" name="recipesData" />
      <input type="submit" value="upload" />
    </form>
  </body>
</html>