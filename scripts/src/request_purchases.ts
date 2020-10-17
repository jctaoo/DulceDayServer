import * as fs from "fs";
import * as path from "path";
import {AppStoreConnectClient} from "./app_store_connect_client";
import * as bombadil from "@sgarciac/bombadil";
import * as mysql from "mysql";

const __DEV__ = true; // todo

const projectRoot = path.join(__dirname, "..", "..")
const tomlConfigFilePath = path.join(projectRoot, "config.toml");
const privateP8KeyFilePath = path.join(projectRoot, "AuthKey_S5FZB2J5J3.p8");

let config;
const configRaw = fs.readFileSync(tomlConfigFilePath).toString();
const privateKey = fs.readFileSync(privateP8KeyFilePath).toString();

const reader = new bombadil.TomlReader();
reader.readToml(configRaw);
config = reader.result;

let dataBaseConfig = __DEV__ ? config.dataBaseConfigs.development : config.dataBaseConfigs.production;
const connection = mysql.createConnection({
    host: dataBaseConfig.Host,
    port: dataBaseConfig.Port,
    user: dataBaseConfig.User,
    password: dataBaseConfig.Password,
    database: dataBaseConfig.Collection,
    insecureAuth: true,
});
connection.connect();

const client = new AppStoreConnectClient(
    privateKey,
    config.appStoreConnect.apiKey,
    config.appStoreConnect.issId
);
const app = client.withAppId(config.appStoreConnect.appId);
app.getInAppPurchaseItems().then(items => {
    if (!!items) {
        // todo 区分更改和添加，避免删除，同时实现 createdAt 等字段
        const valuesStr = items.map(item => {
            return `("${item.attributes.productId}", "${item.attributes.referenceName}", ${item.attributes.inAppPurchaseType})`;
        }).join(", \n")
        const query = `
            INSERT INTO purchase_items
                (internal_identifier, internal_name, type)
            VALUES 
                ${valuesStr};
        `;
        connection.query(query, (err, result, fields) => {
            if (err == null) {
                connection.end();
            } else {
                console.error("写入数据库时发生错误")
            }
        });
    }
});