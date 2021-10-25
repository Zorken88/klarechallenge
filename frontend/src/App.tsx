import React, { useEffect, useState } from 'react';
import './App.css';
import { useProduct } from './hooks/useProduct';
import CollapsibleTable from './components/Table/CollapsibleTable';
import { Button, TextField, Typography } from '@mui/material';
import { Product } from './interfaces/product';

function App() {
  const { products, product, getProducts, createProduct, deleteProduct, updateProduct } = useProduct()
  const [name, setName] = useState('')
  const [price, setPrice] = useState('')


  useEffect(() => {
    getProducts()
  }, [product])

  const handleNameChange = (e: any) => {
    setName(e.currentTarget.value)
  }

  const handlePriceChange = (e: any) => {
    setPrice(e.currentTarget.value)
  }

  const handleClick = (e: any) => {

    let newProduct: Product = {
      name: name,
      price: Number.parseFloat(price),
      ID: 0,
      Prices: [],
      CreatedAt: new Date(),
      UpdatedAt: new Date(),
      DeletedAt: null
    }
    createProduct(newProduct)
  }

  const handleDelete = (id: number) => {
    deleteProduct(id)
  }

  const handleUpdate = (id: number, product: Product) => {
    product.name = name
    product.price = Number.parseFloat(price)
    updateProduct(id, product)
  }

  return (
    <div className="App">
      <div style={{ margin: 10, justifyContent: 'center', display: 'flex', flexDirection: 'column' }}>
        <div style={{ flex: 1, alignItems: 'center', justifyContent: 'center' }}>
          <Typography>Insert New Product</Typography>
          <div style={{ alignItems: 'space-between', padding: '20px' }}>
            <TextField id='txfName' label='Name' variant='standard' onChange={handleNameChange} />
            <TextField id='txfPrice' label='Price' variant='standard' onChange={handlePriceChange} type='number' />
            <Button style={{ borderColor: 'blue', borderRadius: 5 }} onClick={handleClick} variant='outlined'>
              <span>Insert</span>
            </Button>
          </div>
        </div>
        <div style={{ flex: 1, marginTop: '20px' }}>
          <CollapsibleTable data={products} handleDelete={handleDelete}
            handleNameChange={handleNameChange}
            handlePriceChange={handlePriceChange}
            handleUpdate={handleUpdate} />
        </div>
      </div>
    </div>
  );
}

export default App;
