#!/usr/bin/env node
const https = require("https");
const fs = require("fs");
const path = require("path");
const { execSync } = require("child_process");

const VERSION = require("../package.json").version;
const REPO = "farquaad/deviantart-mcp";
const BIN_DIR = path.join(__dirname, "..", "bin");
const MARKER = path.join(BIN_DIR, ".installed-version");

// Skip if already installed at current version
if (
  fs.existsSync(MARKER) &&
  fs.readFileSync(MARKER, "utf8").trim() === VERSION
) {
  process.exit(0);
}

const platform = process.platform;
const arch = process.arch;

const osMap = { linux: "linux", darwin: "darwin", win32: "windows" };
const archMap = { x64: "amd64", arm64: "arm64" };

const goos = osMap[platform];
const goarch = archMap[arch];

if (!goos || !goarch) {
  console.error(`Unsupported platform/arch: ${platform}/${arch}`);
  process.exit(1);
}

const ext = goos === "windows" ? ".exe" : "";
const archiveExt = goos === "windows" ? ".zip" : ".tar.gz";
const assetName = `deviantart-mcp_${goos}_${goarch}${archiveExt}`;
const downloadUrl = `https://github.com/${REPO}/releases/download/v${VERSION}/${assetName}`;

console.log(`Downloading deviantart-mcp v${VERSION} for ${goos}/${goarch}...`);
console.log(downloadUrl);

const tmpArchive = path.join(BIN_DIR, `deviantart-mcp${archiveExt}`);

function download(url, dest) {
  return new Promise((resolve, reject) => {
    const file = fs.createWriteStream(dest);
    https
      .get(url, { headers: { "User-Agent": "npm-install" } }, (res) => {
        if (res.statusCode === 302 || res.statusCode === 301) {
          file.close();
          fs.unlinkSync(dest);
          download(res.headers.location, dest).then(resolve).catch(reject);
          return;
        }
        if (res.statusCode !== 200) {
          file.close();
          fs.unlinkSync(dest);
          reject(new Error(`HTTP ${res.statusCode}: ${res.statusMessage}`));
          return;
        }
        res.pipe(file);
        file.on("finish", () => {
          file.close();
          resolve();
        });
      })
      .on("error", reject);
  });
}

async function main() {
  fs.mkdirSync(BIN_DIR, { recursive: true });

  await download(downloadUrl, tmpArchive);

  if (goos === "windows") {
    // Extract zip
    const AdmZip = require("adm-zip");
    const zip = new AdmZip(tmpArchive);
    zip.extractAllTo(BIN_DIR, true);
    fs.unlinkSync(tmpArchive);
  } else {
    // Extract tar.gz
    execSync(`tar -xzf "${tmpArchive}" -C "${BIN_DIR}"`, { stdio: "inherit" });
    fs.unlinkSync(tmpArchive);
  }

  const binPath = path.join(BIN_DIR, `deviantart-mcp${ext}`);
  if (!fs.existsSync(binPath)) {
    console.error("Binary not found after extraction");
    process.exit(1);
  }

  if (goos !== "windows") {
    fs.chmodSync(binPath, 0o755);
  }

  fs.writeFileSync(MARKER, VERSION);
  console.log(`deviantart-mcp v${VERSION} installed successfully`);
}

main().catch((err) => {
  console.error("Install failed:", err.message);
  process.exit(1);
});
