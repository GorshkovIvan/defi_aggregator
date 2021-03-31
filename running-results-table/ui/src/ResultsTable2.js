import React from 'react';
import { Table, Header, Segment, Label } from 'semantic-ui-react'

export default function ResultsTable2({currencyoutputtable}) {
    const rows = currencyoutputtable.map(((result, index) => {
        let color = 'blue';

        return (
            <Table.Row key={ index }>
                <Table.Cell>
                    <Label ribbon color={color}>{ index + 1 }</Label>
                </Table.Cell>
                <Table.Cell>{ result.backend_pair }</Table.Cell>
                <Table.Cell>{ result.backend_poolsize }</Table.Cell>
                <Table.Cell>{ result.backend_volume }</Table.Cell>
                <Table.Cell>{ result.backend_yield }</Table.Cell>
                <Table.Cell>{ result.pool_source }</Table.Cell>
                <Table.Cell>{ result.volatility }</Table.Cell>
                <Table.Cell>{ result.ROIestimate }</Table.Cell>
                <Table.Cell>{ result.ROIhist }</Table.Cell>
            </Table.Row>
        );
    }));
    return (
        <div className="ui container">
            <Segment>

                <Table>
                    <Table.Header>
                        <Table.Row>
                            <Table.HeaderCell>Ranking</Table.HeaderCell>
                            <Table.HeaderCell>Currency Pair</Table.HeaderCell>
                            <Table.HeaderCell>Pool Size</Table.HeaderCell>
                            <Table.HeaderCell>Pool Trading Volume</Table.HeaderCell>
                            <Table.HeaderCell>Interest Rate</Table.HeaderCell>
                            <Table.HeaderCell>Pool</Table.HeaderCell>
                            <Table.HeaderCell>Historical Volatility</Table.HeaderCell>
                            <Table.HeaderCell>ROI Est Raw</Table.HeaderCell>
                            <Table.HeaderCell>ROI Hist</Table.HeaderCell>
                        </Table.Row>
                    </Table.Header>
                    <Table.Body>
                        { rows }
                    </Table.Body>
                </Table>
            </Segment>
        </div>
    );
}


       /*
        if (index === 0) {
            color='blue';
        } else if (index === 1) {
            color='blue';
        } else if (index === 2) {
            color='blue';
        }

             <Table.Cell>{ result.max_amount }</Table.Cell>
                <Table.Cell>{ result.raw_yield }</Table.Cell>
                <Table.Cell>{ result.normalized_yield }</Table.Cell>

        */