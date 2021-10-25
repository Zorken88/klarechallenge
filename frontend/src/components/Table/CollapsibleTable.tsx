import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import Paper from '@mui/material/Paper';
import { Product } from '../../interfaces/product';
import { Row } from './Row';

const CollapsibleTable = (props: 
    { 
        data: Product[], 
        handleDelete: (id: number) => void,
        handleNameChange: (e: any) => void,
        handlePriceChange : (e: any) => void,
        handleUpdate : (id: number, product: Product) => void 
    }) => {
    const { data, handleDelete, handleNameChange, handlePriceChange, handleUpdate } = props
    return (
        <TableContainer component={Paper}>
            <Table aria-label="collapsible table">
                <TableHead>
                    <TableRow>
                        <TableCell />
                        <TableCell>ID</TableCell>
                        <TableCell>Name</TableCell>
                        <TableCell>Price</TableCell>
                        <TableCell>CreatedAt</TableCell>
                        <TableCell>UpdatedAt</TableCell>
                        <TableCell>DeletedAt</TableCell>
                        <TableCell>Actions</TableCell>
                    </TableRow>
                </TableHead>
                <TableBody>
                    {data.map((row) => (
                        <Row key={row.ID} row={row} handleDelete={handleDelete} 
                        handleNameChange={handleNameChange} 
                        handlePriceChange={handlePriceChange}
                        handleUpdate={handleUpdate}/>
                    ))}
                </TableBody>
            </Table>
        </TableContainer>
    )
}

export default CollapsibleTable