import * as jwt from "jsonwebtoken";
import fetch, {Headers, Response} from "node-fetch";

export class AppStoreConnectClient {
    readonly #privateKey: string
    readonly #apiKey: string
    readonly #issuerId: string

    constructor(privateKey: string, apiKey: string, issuerId: string) {
        this.#privateKey = privateKey;
        this.#apiKey = apiKey;
        this.#issuerId = issuerId;
    }

    private getAccessToken(): string {
        const now = Math.round((new Date()).getTime() / 1000);
        let expiresIn = 1200;
        const payload = {
            'iss': this.#issuerId,
            'exp': now + expiresIn,
            'aud': 'appstoreconnect-v1'
        };
        const signOptions: jwt.SignOptions = {
            algorithm: "ES256",
            header: {
                'alg': 'ES256',
                'kid': this.#apiKey,
                'typ': 'JWT'
            }
        };
        return jwt.sign(
            payload,
            this.#privateKey,
            signOptions
        );
    }

    public withAppId(appId: string): AppStoreConnectApp {
        return new AppStoreConnectApp(
            appId,
            () => {
                return this.getAccessToken();
            }
        );
    }
}

class AppStoreConnectApp {
    readonly #appId: string
    readonly #getAccessToken: () => string

    constructor(appId: string, getAccessToken: () => string) {
        this.#appId = appId;
        this.#getAccessToken = getAccessToken;
    }

    private generateHeader(): Headers {
        const headers = new Headers();
        headers.set("Authorization", "Bearer " + this.#getAccessToken());
        return headers;
    }

    async getInAppPurchaseItems(): Promise<[InAppPurchase] | null> {
        let rsp: Response;
        try {
            rsp = await fetch(
                `https://api.appstoreconnect.apple.com/v1/apps/${this.#appId}/inAppPurchases`,
                { headers: this.generateHeader() }
            );
            if (rsp.status === 200) {
                const json = await rsp.json();
                if (!!json.data && Array.isArray(json.data)) {
                    return json.data.map((item: InAppPurchase) => {
                        return {
                            ...item,
                            attributes: {
                                ...item.attributes,
                                inAppPurchaseType: InAppPurchaseAttributesType[item.attributes.inAppPurchaseType]
                            }
                        }
                    })
                }
            }
        } catch (e) {
            // 网络错误
            return null;
        }
        return null;
    }
}

interface InAppPurchase {
    attributes: InAppPurchaseAttributes
    id: string
    type: "inAppPurchases"
    links: InAppPurchaseLinks
}

interface InAppPurchaseLinks {
    self: string
}

interface InAppPurchaseAttributes {
    productId: string
    inAppPurchaseType: InAppPurchaseAttributesType
    referenceName: string
    // todo state
    // https://developer.apple.com/documentation/appstoreconnectapi/inapppurchase/attributes
}

enum InAppPurchaseAttributesType {
    CONSUMABLE,
    NON_CONSUMABLE,
    NON_RENEWING_SUBSCRIPTION,
    AUTOMATICALLY_RENEWABLE_SUBSCRIPTION,
    FREE_SUBSCRIPTION,
}
