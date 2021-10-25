import { Product } from "./product";

export interface Response {
    error: boolean,
    msg: string,
    products: Product | Product[]
}