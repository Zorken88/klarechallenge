import { Http } from '@mui/icons-material'
import axios, { Method } from 'axios'
import { useState } from 'react'
import { Product } from '../interfaces/product'
import { Response } from '../interfaces/response'


export const useProduct = () => {
    const [products, setProducts] = useState<Product[]>([])
    const [product, setProduct] = useState<Product>(products[0])
    const [loading, setLoading] = useState(false)
    const [error, setError] = useState('')

    const baseUrl = '/api/v1/products'

    // Abort
    const controller = new AbortController()
    const signal = controller.signal
    const cancelTokenSource = axios.CancelToken.source()

    const callApi = (endpoint: string, method: Method, data?: any) => {
        if (signal.aborted) {
            return Promise.reject()
        }

        return new Promise<void>((resolve, reject) => {
            setError('')
            setLoading(true)

            const options = {
                url: `${baseUrl}${endpoint}`,
                method: method,
                headers: {
                    Accept: "application/json",
                    "Content-Type": "application/json",
                },
                data: data,
                cancelToken: cancelTokenSource.token,
            }

            axios(options)
                .then((response: any) => {
                    if (response.data) {
                        const res: Response = response.data
                        switch (res.products.constructor.name) {
                            case 'Array':
                                setProducts(res.products as Product[])
                                break
                            case 'Object':
                                setProduct(res.products as Product)
                                break
                            default:
                                break
                        }
                    }
                    if (response.status == 204) {
                        console.log('pase')
                        setProduct({ 
                            name: '',
                            price: 0,
                            CreatedAt: new Date(),
                            UpdatedAt: new Date(),
                            DeletedAt: null,
                            ID: -1,
                            Prices: []
                        })
                    }
                    resolve()
                })
                .catch((response) => {
                    console.error(response)
                    setError('')
                    reject()
                })
                .finally(() => {
                    setLoading(false)
                })
        })

    }

    // Get list of products
    const getProducts = () => {
        callApi("", "GET")
    }

    // Get product by id
    const getProduct = (id: number) => {
        callApi(`/${id}`, "GET")
    }

    // Get history prices by product id
    const getPrices = (id: number) => {
        callApi(`/${id}/prices`, "GET")
    }

    // Create new product
    const createProduct = (product: Product) => {
        callApi(``, "POST", product)
    }

    // Update product by id
    const updateProduct = (id: number, product: Product) => {
        callApi(`/${id}`, "PATCH", product)
    }

    // Delete product by id
    const deleteProduct = (id: number) => {
        callApi(`/${id}`, "DELETE")
    }

    const cancelRequests = () => {
        controller.abort()
        cancelTokenSource.cancel()
    }

    return {
        getProducts, getProduct, getPrices, createProduct,
        updateProduct, deleteProduct, cancelRequests, loading, 
        error, products, product
    }
}