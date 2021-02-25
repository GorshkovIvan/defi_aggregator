import React from 'react';
import { Table, Header, Segment, Label } from 'semantic-ui-react'
export default function ResultsTable({results}) {
    const rows = results.map(((result, index) => {
        let color;
        if (index === 0) {
            color='yellow';
        } else if (index === 1) {
            color='grey';
        } else if (index === 2) {
            color='orange';
        }
        return (
            <Table.Row key={ index }>
                <Table.Cell>
                    <Label ribbon color={color}>{ index + 1 }</Label>
                </Table.Cell>
                <Table.Cell>{ result.token }</Table.Cell>
                <Table.Cell>{ "Uniswap" }</Table.Cell>
                <Table.Cell>{ 1 }</Table.Cell>
                <Table.Cell>{ result.amount/100 }</Table.Cell>
            </Table.Row>
        );
    }));
    return (
        <div className="ui container">
            <Segment>
                <Header>Optimal Portfolio</Header>
                <Table striped>
                    <Table.Header>
                        <Table.Row>
                            <Table.HeaderCell>Ranking</Table.HeaderCell>
                            <Table.HeaderCell>Pair/Token</Table.HeaderCell>
                            <Table.HeaderCell>Pool</Table.HeaderCell>
                            <Table.HeaderCell>ROI Estimate</Table.HeaderCell>
                            <Table.HeaderCell>% Portfolio</Table.HeaderCell>
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