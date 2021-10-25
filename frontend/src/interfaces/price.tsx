export interface Price {
    ID: number,
    CreatedAt: Date,
    UpdatedAt: Date,
    DeletedAt: Date | null,
    value: number,
    productId: number  
}