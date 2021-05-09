import React from 'react';
import { Form, Table, Segment, Label, Popup } from 'semantic-ui-react'
// currencoutputtable needed for testing, DONT delete
/*
var currencyoutputtable = [
    {
        backend_pair: "ETH/DAI",
        backend_poolsize: 420,
        backend_volume: 420,
        backend_yield: 460,
        pool_source: "Uniswap",
        volatility: 123,
        ROIestimate: 10,
        ROIvoladjest: 100,
        ROIhist: 100,
    },
    {
        backend_pair: "ETH/DAI",
        backend_poolsize: 420,
        backend_volume: 420,
        backend_yield: 460,
        pool_source: "Uniswap",
        volatility: 123,
        ROIestimate: 10,
        ROIvoladjest: 100,
        ROIhist: 100,
    },
    {
        backend_pair: "ETH/DAI",
        backend_poolsize: 420,
        backend_volume: 420,
        backend_yield: 460,
        pool_source: "Uniswap",
        volatility: 123,
        ROIestimate: 10,
        ROIvoladjest: 100,
        ROIhist: 100,
    },
    {
        backend_pair: "ETH/DAI",
        backend_poolsize: 420,
        backend_volume: 420,
        backend_yield: 460,
        pool_source: "Uniswap",
        volatility: 123,
        ROIestimate: 10,
        ROIvoladjest: 100,
        ROIhist: 100,
    }
]*/
//export default function ResultsTable2() {
export default function ResultsTable2({currencyoutputtable}) {
    const rows = currencyoutputtable.map(((result, index) => {
        let color;
        if (index === 0 || index === 1 || index === 2) {
            color='blue';
        } else {
            color='grey';
        }
        return (
            <Table.Row key={ index }>
                <td><Label class="ui horizontal label" color={color}>{ index + 1 }</Label></td>
                <td>{ result.backend_pair }</td>
                <td class="right aligned">{ result.backend_poolsize }</td>
                <td class="right aligned">{ result.backend_volume }</td>
                <td class="right aligned">{ result.backend_yield }</td>
                <td>{ result.pool_source }</td>
                <td class="right aligned">{ result.volatility }</td>
                <td class="right aligned">{ result.ROIestimate }</td>
                <td class="right aligned">{ result.ROIvoladjest}</td>
                <td class="right aligned">{ result.ROIhist }</td>
            </Table.Row>
        );
    }));
    return (
        <div className="ui container">
            <Segment>
                <h4>Recommended Liquidity Pools</h4>
                <Form className="tableButton">
                    <div class="ui animated fade button">
                        <div class="visible content">Refresh</div>
                        <div class="hidden content">
                            <i class="refresh icon"></i>
                        </div>
                    </div>
                </Form>
                <div class="ui basic table">
                    <Table.Header>
                        <Table.Row>
                            <Table.HeaderCell><h3 className="headerTitle">Ranking</h3></Table.HeaderCell>
                            <Table.HeaderCell><h3 className="headerTitle">Currency Pair</h3></Table.HeaderCell>
                            <Table.HeaderCell><h3 className="headerTitle">Pool Size</h3></Table.HeaderCell>
                            <Table.HeaderCell><h3 className="headerTitle">Pool Trading Volume</h3></Table.HeaderCell>
                            <Table.HeaderCell><h3 className="headerTitle">Interest Rate</h3></Table.HeaderCell>
                            <Table.HeaderCell><h3 className="headerTitle">Pool</h3></Table.HeaderCell>
                            <Table.HeaderCell><h3 className="headerTitle">Historical Volatility</h3></Table.HeaderCell>
                            <Table.HeaderCell><h3 className="headerTitle">ROI Est Raw</h3><Popup header="Title" content="Explanation goes here" position="top center" trigger={<i class="info circle icon"></i>} /></Table.HeaderCell>
                            <Table.HeaderCell><h3 className="headerTitle">ROI Est Vol-Adj (Sharpe Ratio)</h3><Popup header="Title" content="Explanation goes here" position="top center" trigger={<i class="info circle icon"></i>} /></Table.HeaderCell>
                            <Table.HeaderCell><h3 className="headerTitle">ROI Hist</h3><Popup header="Title" content="Explanation goes here" position="top center" trigger={<i class="info circle icon"></i>} /></Table.HeaderCell>
                        </Table.Row>
                    </Table.Header>
                    <Table.Body>
                        { rows }
                    </Table.Body>
                </div>
            </Segment>
        </div>
    );
}

/*
    <Table.Cell>{ result.max_amount }</Table.Cell>
    <Table.Cell>{ result.raw_yield }</Table.Cell>
    <Table.Cell>{ result.normalized_yield }</Table.Cell>
*/
