<?php

declare(strict_types=1);

/**
 * PHP Dev Container - Example Application
 *
 * This is a simple example to verify your PHP development environment is working.
 */

// Display PHP info in a nice format
function displaySystemInfo(): void
{
    echo "<div style='font-family: Arial, sans-serif; max-width: 800px; margin: 50px auto; padding: 20px;'>";
    echo "<h1>PHP Dev Container</h1>";
    echo "<p style='color: green; font-size: 18px;'>Your PHP development environment is ready!</p>";

    echo "<h2>Environment Information</h2>";
    echo "<table style='border-collapse: collapse; width: 100%;'>";
    echo "<tr style='background: #f5f5f5;'><td style='padding: 10px; border: 1px solid #ddd;'><strong>PHP Version</strong></td><td style='padding: 10px; border: 1px solid #ddd;'>" . PHP_VERSION . "</td></tr>";
    echo "<tr><td style='padding: 10px; border: 1px solid #ddd;'><strong>Server Software</strong></td><td style='padding: 10px; border: 1px solid #ddd;'>" . ($_SERVER['SERVER_SOFTWARE'] ?? 'N/A') . "</td></tr>";
    echo "<tr style='background: #f5f5f5;'><td style='padding: 10px; border: 1px solid #ddd;'><strong>Document Root</strong></td><td style='padding: 10px; border: 1px solid #ddd;'>" . ($_SERVER['DOCUMENT_ROOT'] ?? 'N/A') . "</td></tr>";
    echo "</table>";

    echo "<h2>Installed Extensions</h2>";
    $extensions = get_loaded_extensions();
    sort($extensions);
    echo "<div style='display: flex; flex-wrap: wrap; gap: 10px;'>";
    foreach ($extensions as $ext) {
        echo "<span style='background: #e0e0e0; padding: 5px 10px; border-radius: 4px; font-size: 12px;'>$ext</span>";
    }
    echo "</div>";

    echo "<h2>Quick Links</h2>";
    echo "<ul>";
    echo "<li><a href='?phpinfo=1'>View Full PHP Info</a></li>";
    echo "</ul>";

    echo "<h2>Next Steps</h2>";
    echo "<ol>";
    echo "<li>Edit files in the <code>src/</code> directory</li>";
    echo "<li>Use Composer to manage dependencies</li>";
    echo "<li>Configure Xdebug in VS Code for debugging</li>";
    echo "</ol>";

    echo "</div>";
}

// Show full phpinfo if requested
if (isset($_GET['phpinfo'])) {
    phpinfo();
    exit;
}

displaySystemInfo();
