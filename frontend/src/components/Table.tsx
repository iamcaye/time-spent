import React from 'react';
import type { IColumn } from '../models/IColumn';
import { ChevronRightIcon } from './icons/ChevronRightIcon';
import { ChevronLeftIcon } from './icons/ChevronLeftIcon';

export function Table(props: { data: any, columns: IColumn[], rows?: number, paginator?: boolean }) {
	const rawData = props.data;
	const { columns, rows, paginator } = props;
	const [data, setData] = React.useState([]);
	const [page, setPage] = React.useState(1);
	const [pageSize, setPageSize] = React.useState(rows || 5);

	React.useEffect(() => {
		const start = (page - 1) * pageSize;
		const end = start + pageSize;
		console.log(rawData.slice(start, end));
		setData(rawData.slice(start, end));
	}, [page, pageSize]);

	return (
		<table className='w-full table-fixed shadow-sm rounded-sm divide-y divide-gray-200'>
			<thead className='bg-gray-200 text-black dark:bg-gray-500 dark:text-white'>
				<tr>
					{columns.map((col: IColumn, index: number) => (
						<th className=' border border-solid border-black' key={index}>{col.header}</th>
					))}
				</tr>
			</thead>
			<tbody>
				{data.map((row, index) => (
					<tr key={index} className=''>
						{columns.map((col: IColumn, index: number) => (
							<td
								className='border border-solid border-black text-center'
								key={index}>
								{col.type == 'text' && row[col.field]}
								{col.type == 'date' && new Date(row[col.field]).toLocaleDateString()}
								{col.type == 'number' && Number(row[col.field]).toLocaleString()}
								{col.type == 'currency' && Number(row[col.field]).toLocaleString('es-EU', { style: 'currency', currency: 'EUR' })}
							</td>
						))}
					</tr>
				))}
			</tbody>
			<tfoot>
				<tr>
					{paginator && (
						<td colSpan={columns.length} className='relative border border-solid border-black'>
							<div className='flex justify-center'>
								<span className='flex items-center gap-8'>
									<button className='p-1 m-1 bg-gray-200 rounded-full disabled:opacity-50'
										onClick={() => setPage(page - 1)}
										disabled={page == 1}
									>
										<ChevronLeftIcon />
									</button>
									<span>{page}/{Math.ceil(rawData.length / pageSize)}</span>
									<button className='p-1 m-1 bg-gray-200 rounded-full disabled:opacity-50'
										onClick={() => setPage(page + 1)}
										disabled={data.length < (page - 1) * pageSize}
									>
										<ChevronRightIcon />
									</button>
								</span>
							</div>
							<div className='absolute flex justify-center items-center gap-2 right-2 top-2'>
								<select value={pageSize} onChange={(e) => setPageSize(Number(e.target.value))}>
									<option value={5}>5</option>
									<option value={10}>10</option>
									<option value={20}>20</option>
									<option value={50}>50</option>
								</select>
							</div>
						</td>
					)}
				</tr>
			</tfoot>


		</table >
	);
}

