import React, { Fragment, useCallback, useEffect, useState } from 'react';
import Box from '@mui/material/Box';
import Collapse from '@mui/material/Collapse';
import IconButton from '@mui/material/IconButton';
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import Typography from '@mui/material/Typography';
import KeyboardArrowDownIcon from '@mui/icons-material/KeyboardArrowDown';
import KeyboardArrowUpIcon from '@mui/icons-material/KeyboardArrowUp';
import DeleteIcon from '@mui/icons-material/Delete'
import EditIcon from '@mui/icons-material/Edit'
import { Product } from '../../interfaces/product'
import { useProduct } from '../../hooks/useProduct';
import { TextField } from '@mui/material';

export const Row = (props: 
    {   row: Product, 
        handleDelete: (id: number) => void,
        handleNameChange: (e: any) => void,
        handlePriceChange : (e: any) => void,
        handleUpdate : (id: number, product: Product) => void
    }) => {
    const { row, handleDelete, handleNameChange, handlePriceChange, handleUpdate } = props;
    const [open, setOpen] = useState(false);

    const { getPrices, product } = useProduct()

    useEffect(() => {
        getPrices(row.ID)
    }, [row])

    const handle = () => {
        handleDelete(row.ID)
    }

    const handleModify = () => {
        handleUpdate(row.ID, product)
    }

    return (
        <Fragment>
            <TableRow sx={{ '& > *': { borderBottom: 'unset' } }}>
                <TableCell>
                    <IconButton
                        aria-label="expand row"
                        size="small"
                        onClick={() => {
                            row.Prices = product.Prices
                            setOpen(!open)
                        }}
                    >
                        {open ? <KeyboardArrowUpIcon /> : <KeyboardArrowDownIcon />}
                    </IconButton>
                </TableCell>
                <TableCell >{row.ID}</TableCell>

                <TableCell>
                    <TextField id='txfName' label='Name' variant='standard' onChange={handleNameChange} defaultValue={row.name} />
                </TableCell>
                <TableCell >
                    <TextField id='txfPrice' label='Price' variant='standard' onChange={handlePriceChange} type='number' defaultValue={row.price} />
                </TableCell>

                <TableCell >{row.CreatedAt} </TableCell>
                <TableCell >{row.UpdatedAt}</TableCell>
                <TableCell >{row.DeletedAt}</TableCell>
                <TableCell>
                <span style={{cursor: 'pointer'}} onClick={handleModify}><EditIcon/></span>
                <span style={{cursor: 'pointer'}} onClick={handle}><DeleteIcon/></span>
                </TableCell>
            </TableRow>
            {open &&
                <TableRow>
                    <TableCell style={{ paddingBottom: 0, paddingTop: 0 }} colSpan={6}>
                        <Collapse in={open} timeout="auto" unmountOnExit>
                            <Box sx={{ margin: 1 }}>
                                <Typography variant="h6" gutterBottom component="div">
                                    History
                                </Typography>
                                <Table size="small" aria-label="purchases">
                                    <TableHead>
                                        <TableRow>
                                            <TableCell>ID</TableCell>
                                            <TableCell>Value</TableCell>
                                            <TableCell>ProductId</TableCell>
                                            <TableCell>CreatedAt</TableCell>
                                            <TableCell>UpdatedAt</TableCell>
                                            <TableCell>DeletedAt</TableCell>
                                        </TableRow>
                                    </TableHead>
                                    <TableBody>
                                        {row.Prices.length > 0 && row.Prices.map((pricesRow) => (
                                            <TableRow key={pricesRow.ID}>
                                                <TableCell component="th" scope="row">
                                                    {pricesRow.ID}
                                                </TableCell>
                                                <TableCell>{pricesRow.value}</TableCell>
                                                <TableCell >{pricesRow.productId}</TableCell>
                                                <TableCell >
                                                    {pricesRow.CreatedAt}
                                                </TableCell>
                                                <TableCell >
                                                    {pricesRow.UpdatedAt}
                                                </TableCell>
                                                <TableCell >
                                                    {pricesRow.DeletedAt}
                                                </TableCell>
                                            </TableRow>
                                        ))}
                                    </TableBody>
                                </Table>
                            </Box>
                        </Collapse>
                    </TableCell>
                </TableRow>
            }
        </Fragment>
    );
}
