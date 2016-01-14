<html>
    <body>
        <h1>$FirstName</h1>

        <p>Here's a list of your favorite colors:</p>
        <ul>
        $for _, colorName in FavoriteColors:
            <li>$colorName</li>
        $end
        </ul>
    </body>
</html>