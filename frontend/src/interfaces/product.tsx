import { Price } from "./price";

export interface Product {
    ID: number,
    CreatedAt: Date,
    UpdatedAt: Date,
    DeletedAt: Date | null,
    name: string,
    price: number,
    Prices: Price[]
}
