package com.finsim.xchanger.orders.configuration;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.CommandLineRunner;
import org.springframework.data.domain.Sort;
import org.springframework.data.mongodb.core.DefaultIndexOperations;
import org.springframework.data.mongodb.core.MongoTemplate;
import org.springframework.data.mongodb.core.index.Index;
import org.springframework.data.mongodb.core.index.PartialIndexFilter;
import org.springframework.data.mongodb.core.query.Criteria;
import org.springframework.stereotype.Component;

import com.finsim.xchanger.orders.model.Order;

@Component
public class OrderConfiguration implements CommandLineRunner {
    @Autowired
    private MongoTemplate mongoTemplate;

	@Override
	public void run(String... args) throws Exception {
		System.out.printf("Order collection configuration started%n");

		Index orderBuyIndex = new Index()
            .named("order-buy-index")
            .on("isin", Sort.Direction.ASC)
            .on("price", Sort.Direction.DESC)
            .partial(PartialIndexFilter.of(
                Criteria
                    .where("leftQuantity").gt(0)
                    .and("type").is("BUY")
            ))
        ;
		Index orderSellIndex = new Index()
            .named("order-sell-index")
            .on("isin", Sort.Direction.ASC)
            .on("price", Sort.Direction.ASC)
            .partial(PartialIndexFilter.of(
                Criteria
                    .where("leftQuantity").gt(0)
                    .and("type").is("SELL")
            ))
        ;

        DefaultIndexOperations indexOperations = new DefaultIndexOperations(
            mongoTemplate,
            mongoTemplate.getCollectionName(Order.class),
            getClass()
        );
        indexOperations.createIndex(orderBuyIndex);
        indexOperations.createIndex(orderSellIndex);

		System.out.printf("Order collection configuration ended%n");
	}
}
