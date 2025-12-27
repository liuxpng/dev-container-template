#!/usr/bin/env python3
"""
Python Dev Container - Example Application

This is a simple HTTP server example to verify your Python development environment is working.
"""

import sys
from http.server import HTTPServer, SimpleHTTPRequestHandler
from typing import Any


class DevContainerHandler(SimpleHTTPRequestHandler):
    """Custom HTTP handler for the dev container example."""

    def do_GET(self) -> None:
        """Handle GET requests."""
        if self.path == "/" or self.path == "/index.html":
            self.send_response(200)
            self.send_header("Content-type", "text/html; charset=utf-8")
            self.end_headers()
            self.wfile.write(self.get_html_content().encode())
        elif self.path == "/api/info":
            self.send_response(200)
            self.send_header("Content-type", "application/json")
            self.end_headers()
            self.wfile.write(self.get_json_info().encode())
        else:
            super().do_GET()

    def get_html_content(self) -> str:
        """Generate HTML content for the main page."""
        return f"""
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Python Dev Container</title>
    <style>
        body {{
            font-family: Arial, sans-serif;
            max-width: 800px;
            margin: 50px auto;
            padding: 20px;
            background: #f5f5f5;
        }}
        .container {{
            background: white;
            padding: 30px;
            border-radius: 10px;
            box-shadow: 0 2px 10px rgba(0,0,0,0.1);
        }}
        h1 {{ color: #3776ab; }}
        .success {{ color: green; font-size: 18px; }}
        table {{ border-collapse: collapse; width: 100%; margin: 20px 0; }}
        td {{ padding: 10px; border: 1px solid #ddd; }}
        tr:nth-child(even) {{ background: #f9f9f9; }}
        code {{ background: #e0e0e0; padding: 2px 6px; border-radius: 4px; }}
        .tools {{ display: flex; flex-wrap: wrap; gap: 10px; margin: 20px 0; }}
        .tool {{ background: #3776ab; color: white; padding: 5px 15px; border-radius: 20px; font-size: 14px; }}
    </style>
</head>
<body>
    <div class="container">
        <h1>Python Dev Container</h1>
        <p class="success">Your Python development environment is ready!</p>

        <h2>Environment Information</h2>
        <table>
            <tr><td><strong>Python Version</strong></td><td>{sys.version}</td></tr>
            <tr><td><strong>Platform</strong></td><td>{sys.platform}</td></tr>
            <tr><td><strong>Executable</strong></td><td>{sys.executable}</td></tr>
        </table>

        <h2>Installed Tools</h2>
        <div class="tools">
            <span class="tool">pip</span>
            <span class="tool">poetry</span>
            <span class="tool">uv</span>
            <span class="tool">pytest</span>
            <span class="tool">ruff</span>
            <span class="tool">black</span>
            <span class="tool">mypy</span>
            <span class="tool">ipython</span>
        </div>

        <h2>Quick Links</h2>
        <ul>
            <li><a href="/api/info">View JSON API Info</a></li>
        </ul>

        <h2>Next Steps</h2>
        <ol>
            <li>Edit files in the <code>src/</code> directory</li>
            <li>Use <code>poetry</code> or <code>uv</code> to manage dependencies</li>
            <li>Run tests with <code>pytest</code></li>
            <li>Format code with <code>ruff format</code> or <code>black</code></li>
        </ol>
    </div>
</body>
</html>
"""

    def get_json_info(self) -> str:
        """Generate JSON info about the environment."""
        import json
        import platform

        info = {
            "status": "ok",
            "python": {
                "version": sys.version,
                "platform": sys.platform,
                "implementation": platform.python_implementation(),
            },
            "tools": ["pip", "poetry", "uv", "pytest", "ruff", "black", "mypy", "ipython"],
        }
        return json.dumps(info, indent=2)


def main() -> None:
    """Run the development server."""
    port = 8000
    server_address = ("", port)
    httpd = HTTPServer(server_address, DevContainerHandler)
    print(f"Python Dev Container Server")
    print(f"Serving at http://localhost:{port}")
    print("Press Ctrl+C to stop")
    try:
        httpd.serve_forever()
    except KeyboardInterrupt:
        print("\nServer stopped.")


if __name__ == "__main__":
    main()
